package main

import "fmt"

// 策略模式： 一个类的行为或其算法可以在运行时更改。这种类型的设计模式属于行为型模式
// 策略模式与工厂模式很相似， 他们的区别：
// 		工厂模式是创建模式，在程序运行时，创建对象的类型是已知的;
//		策略模式是行为模式，在程序运行时，是未知的，根据context参数来决定算法或行为

type Eat interface {
	Eat()
}

type FoodRice struct {
}

func (*FoodRice) Eat() {
	fmt.Println("eat rice ...")
}

type FoodNoodle struct {
}

func (*FoodNoodle) Eat() {
	fmt.Println("eat noodle ...")
}

// 创建策略
func getFood(food string) Eat {
	switch food {
	case "rice":
		return &FoodRice{}
	case "noodle":
		return &FoodNoodle{}
	}
	return nil
}

func main() {
	food := getFood("rice")
	food.Eat()

	food = getFood("noodle")
	food.Eat()
}
