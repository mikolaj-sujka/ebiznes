package com.example.plugins

import com.slack.api.Slack
import com.slack.api.methods.SlackApiException
import io.ktor.server.application.*
import io.ktor.server.request.*
import io.ktor.server.response.*
import io.ktor.server.routing.*
import kotlinx.serialization.encodeToString
import kotlinx.serialization.json.Json
import net.dv8tion.jda.api.JDABuilder
import net.dv8tion.jda.api.entities.MessageChannel
import net.dv8tion.jda.api.events.message.MessageReceivedEvent
import net.dv8tion.jda.api.hooks.ListenerAdapter
import javax.security.auth.login.LoginException

fun Application.configureRouting() {
    val dotenv = dotenv {
        ignoreIfMissing = true 
    }

    val slackToken = System.getenv("SLACK_TOKEN")
    val slackChannelName = System.getenv("SLACK_CHANNEL_NAME")
    val slack = Slack.getInstance()

    val discordToken = System.getenv("DISCORD_TOKEN")
    val discordChannelId = System.getenv("DISCORD_CHANNEL_ID")

    val database = Database(ArrayList())

    try {
        val jda = JDABuilder.createDefault(discordToken)
            .addEventListeners(object : ListenerAdapter() {
                override fun onMessageReceived(event: MessageReceivedEvent) {
                    if (event.author.isBot) return
                    val channel = event.channel
                    val content = event.message.contentRaw
                    when {
                        content.startsWith("!categories") -> {
                            val categories = database.categories
                            val message = if (categories.isNotEmpty()) {
                                "List of categories:\n" + categories.joinToString("\n") { it.name }
                            } else {
                                "No categories available."
                            }
                            channel.sendMessage(message).queue()
                        }
                        content.startsWith("!addCategory") -> {
                            val categoryName = content.removePrefix("!addCategory").trim()
                            val exists = database.categories.any { it.name == categoryName }
                            val response = if (exists) {
                                "Category with the name '$categoryName' already exists."
                            } else {
                                val category = Category(categoryName, ArrayList())
                                database.categories.add(category)
                                "Category '$categoryName' added successfully."
                            }
                            channel.sendMessage(response).queue()
                        }
                        content.startsWith("!addItem") -> {
                            val parts = content.removePrefix("!addItem").trim().split(" ", limit = 3)
                            if (parts.size < 3) {
                                channel.sendMessage("Usage: !addItem <itemName> <price> <categoryName>").queue()
                                return
                            }
                            val (itemName, priceStr, categoryName) = parts
                            val price = priceStr.toDoubleOrNull()
                            if (price == null) {
                                channel.sendMessage("Invalid price format. Please enter a valid number.").queue()
                                return
                            }
                            val category = database.categories.find { it.name == categoryName }
                            val response = if (category != null) {
                                val itemExists = category.items.any { it.name == itemName }
                                if (itemExists) {
                                    "Item with the name '$itemName' already exists in category '$categoryName'."
                                } else {
                                    val item = Item(itemName, price, categoryName)
                                    category.items.add(item)
                                    "Item '$itemName' added to category '$categoryName'."
                                }
                            } else {
                                "Category '$categoryName' does not exist."
                            }
                            channel.sendMessage(response).queue()
                        }
                        content.startsWith("!getItems") -> {
                            val categoryName = content.removePrefix("!getItems").trim()
                            val category = database.categories.find { it.name == categoryName }
                            val message = if (category != null && category.items.isNotEmpty()) {
                                "Items in category '$categoryName':\n" +
                                        category.items.joinToString("\n") { "${it.name} - ${it.price}$" }
                            } else {
                                "No items found in category '$categoryName'."
                            }
                            channel.sendMessage(message).queue()
                        }
                    }
                }
            })
            .build()
        jda.awaitReady()
    } catch (e: LoginException) {
        println("Failed to login to Discord: ${e.message}")
    }

    routing {
        // Slack integration
        post("/getCategories") {
            val textToSend = if (database.categories.isNotEmpty()) {
                "List of categories:\n" + database.categories.joinToString("\n") { it.name }
            } else {
                "Categories do not exist."
            }
            try {
                slack.methods(slackToken).chatPostMessage { it.channel(slackChannelName).text(textToSend) }
            } catch (e: SlackApiException) {
                println("Error posting to Slack: ${e.message}")
            }
            call.respondText(Json.encodeToString(database.categories))
        }

        post("/getItemsByCategory") {
            val categoryName = call.receiveParameters()["text"].toString()
            val category = database.categories.find { it.name == categoryName }
            val responseMessage = if (category != null && category.items.isNotEmpty()) {
                "List of items:\n" + category.items.joinToString("\n") { "${it.name} ${it.price}$" }
            } else {
                "Category with given name does not exist or has no items."
            }
            try {
                slack.methods(slackToken).chatPostMessage { it.channel(slackChannelName).text(responseMessage) }
            } catch (e: SlackApiException) {
                println("Error posting to Slack: ${e.message}")
            }
            call.respondText(responseMessage)
        }

        post("/addCategory") {
            val categoryName = call.parameters["categoryName"].toString()
            val exists = database.categories.any { it.name == categoryName }
            val response = if (exists) {
                "Category with given name already exists: $categoryName"
            } else {
                val category = Category(categoryName, ArrayList())
                database.categories.add(category)
                Json.encodeToString(category)
            }
            call.respondText(response)
        }

        post("/addItem") {
            val itemName = call.parameters["itemName"].toString()
            val price = call.parameters["price"]?.toDoubleOrNull() ?: 0.0
            val categoryName = call.parameters["categoryName"].toString()
            val category = database.categories.find { it.name == categoryName }
            val response = if (category != null) {
                val itemExists = category.items.any { it.name == itemName }
                if (itemExists) {
                    "Item with given name already exists: $itemName"
                } else {
                    val item = Item(itemName, price, categoryName)
                    category.items.add(item)
                    Json.encodeToString(item)
                }
            } else {
                "Category for given item does not exist: $categoryName"
            }
            call.respondText(response)
        }

        delete("/deleteCategory") {
            val categoryName = call.parameters["categoryName"].toString()
            val category = database.categories.find { it.name == categoryName }
            val response = if (category != null) {
                database.categories.remove(category)
                Json.encodeToString(category)
            } else {
                "Category with given name does not exist: $categoryName"
            }
            call.respondText(response)
        }

        delete("/deleteItem") {
            val itemName = call.parameters["itemName"].toString()
            val categoryName = call.parameters["categoryName"].toString()
            val category = database.categories.find { it.name == categoryName }
            val response = if (category != null) {
                val item = category.items.find { it.name == itemName }
                if (item != null) {
                    category.items.remove(item)
                    Json.encodeToString(item)
                } else {
                    "Item with given name does not exist: $itemName"
                }
            } else {
                "Category with given name does not exist: $categoryName"
            }
            call.respondText(response)
        }
    }
}
