package com.example.plugins

import kotlinx.serialization.Serializable

@Serializable
class Item(
    val name: String,
    val price: Double,
    val categoryName: String,
)