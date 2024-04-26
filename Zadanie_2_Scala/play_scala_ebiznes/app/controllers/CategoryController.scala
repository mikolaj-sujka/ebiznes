package controllers

import infrastructure.Database
import models._
import play.api.libs.json._
import play.api.mvc._
import javax.inject._

@Singleton
class CategoryController @Inject()(val controllerComponents: ControllerComponents) extends BaseController {

  implicit val jsonCategory = Json.format[Category]
  def getCategories() = Action { implicit request: Request[AnyContent] =>
    Ok(Json.toJson(Database.categories))
  }

  def getCategoryById(categoryId: Long) = Action { implicit request: Request[AnyContent] =>
    val categoryFromDB = Database.categories.filter(category => category.id == categoryId)
    if (categoryFromDB.isEmpty) {
      NotFound("Category not found")
    }
    else {
      Ok(Json.toJson(categoryFromDB.head))
    }
  }

  def addCategory(categoryId: Long, name: String) = Action { implicit request: Request[AnyContent] =>
    val categoryExists = Database.categories.exists(category => category.id == categoryId)
    if (categoryExists) {
      NotAcceptable("Category with given id already exists")
    }
    else {
      val category = Category(categoryId, name)
      Database.categories += category
      Ok(Json.toJson(category))
    }
  }

  def updateCategory(categoryId: Long, name: String) = Action { implicit request: Request[AnyContent] =>
    val categoryFromDB = Database.categories.filter(category => category.id == categoryId)
    if (categoryFromDB.isEmpty) {
      NotFound("Category not found")
    }
    else {
      Database.categories -= categoryFromDB.head
      val updatedCategory = Category(categoryFromDB.head.id, name)
      Database.categories += updatedCategory
      Ok(Json.toJson(updatedCategory))
    }
  }

  def deleteCategory(categoryId: Long) = Action { implicit request: Request[AnyContent] =>
    val categoryFromDB = Database.categories.filter(category => category.id == categoryId)
    if (categoryFromDB.isEmpty) {
      NotFound("Category not found")
    }
    else {
      Database.categories -= categoryFromDB.head
      Ok(Json.toJson(categoryFromDB.head))
    }
  }
}
