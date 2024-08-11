package com.example.plugins

import kotlinx.serialization.Serializable

@Serializable
class Category(
    val name: String,
    val items: ArrayList<Item>,
)