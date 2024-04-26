package infrastructure

import models._

import scala.collection.mutable

object Database {
  val products = new mutable.ListBuffer[Product]()
  val orders = new mutable.ListBuffer[Order]()
  val categories = new mutable.ListBuffer[Category]()
}
