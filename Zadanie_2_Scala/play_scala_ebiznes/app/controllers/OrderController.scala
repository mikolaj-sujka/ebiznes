package controllers

import infrastructure.Database
import models._
import play.api.libs.json._
import play.api.mvc._

import javax.inject._

@Singleton
class OrderController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {
  implicit val productFormat = Json.format[Product]
  implicit val orderFormat = Json.format[Order]

  def createOrder(productId: Long) = Action {
    val product = Database.products.find(_.id == productId)
    product match {
      case Some(p) =>
        val newOrder = Order(Database.orders.size.toLong + 1, List(p))
        Database.orders += newOrder
        Created(Json.toJson(newOrder))
      case None =>
        NotFound("Product not found")
    }
  }

  def addToOrder(orderId: Long, productId: Long) = Action {
    Database.orders.find(_.id == orderId).flatMap { order =>
      Database.products.find(_.id == productId).map { product =>
        val updatedOrder = order.copy(products = order.products :+ product)
        Database.orders -= order
        Database.orders += updatedOrder
        Ok(Json.toJson(updatedOrder))
      }
    }.getOrElse(NotFound("Order or Product not found"))
  }

  def removeFromOrder(orderId: Long, productId: Long) = Action {
    Database.orders.find(_.id == orderId).flatMap { order =>
      val updatedProducts = order.products.filterNot(_.id == productId)
      if (updatedProducts.size == order.products.size) None
      else {
        val updatedOrder = order.copy(products = updatedProducts)
        Database.orders -= order
        Database.orders += updatedOrder
        Some(Ok(Json.toJson(updatedOrder)))
      }
    }.getOrElse(NotFound("Order or Product not found"))
  }

  def getOrder(orderId: Long) = Action {
    Database.orders.find(_.id == orderId) match {
      case Some(order) => Ok(Json.toJson(order))
      case None => NotFound("Order not found")
    }
  }
}
