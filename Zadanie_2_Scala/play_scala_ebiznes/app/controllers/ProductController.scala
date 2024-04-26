package controllers

import infrastructure.Database
import play.api.libs.json._
import play.api.mvc._
import models._
import javax.inject._

@Singleton
class ProductController @Inject()(val controllerComponents: ControllerComponents) extends BaseController{
  implicit val jsonProduct = Json.format[Product]

  def getProducts() = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(Database.products))
  }

  def getProductById(productId: Long) = Action { implicit request: Request[AnyContent] =>
    val productDb = Database.products.find(_.id == productId)
    if (productDb.isEmpty) {
      NotFound("Product not found")
    }
    else {
      Ok(Json.toJson(productDb))
    }
  }

  def addProduct(productId: Long, name: String, categoryId: Long) = Action { implicit request: Request[AnyContent] =>
    val productExists = Database.products.exists(product => product.id == productId)
    val categoryExists = Database.categories.exists(category => category.id == categoryId)
    if (productExists) {
      NotAcceptable("Product with given id already exists")
    }
    else if (!categoryExists) {
      NotAcceptable("Category with given id does not exist")
    }
    else {
      val product = Product(productId, name, categoryId)
      Database.products += product
      Ok(Json.toJson(product))
    }
  }

  def updateProduct(productId: Long, name: String, categoryId: Long) = Action { implicit request: Request[AnyContent] =>
    val productDb = Database.products.filter(product => product.id == productId)
    val categoryExists = Database.categories.exists(category => category.id == categoryId)
    if (productDb.isEmpty) {
      NotFound("Product not found")
    }
    else if (!categoryExists) {
      NotAcceptable("Category with given id does not exist")
    }
    else {
      Database.products -= productDb.head
      val updatedProduct = Product(productDb.head.id, name, categoryId)
      Database.products += updatedProduct
      Ok(Json.toJson(updatedProduct))
    }
  }

  def deleteProduct(productId: Long) = Action { implicit request: Request[AnyContent] =>
    val productDb = Database.products.filter(product => product.id == productId)
    if (productDb.isEmpty) {
      NotFound("Product not found")
    }
    else {
      Database.products -= productDb.head
      Ok(Json.toJson(productDb))
    }
  }
}
