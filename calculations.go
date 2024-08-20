package main

import "math"

func drawPolygons() ([]polygon, []polygon) {
	dartTransforms, kiteTransforms := iterate([]transformation{{0, coordinate{0, 0}, 0}}, []transformation{})
	resultDart, resultKite := []polygon{}, []polygon{}
	for _, trans := range dartTransforms {
		resultDart = append(resultDart, dart.applyTransformation(trans))
	}

	for _, trans := range kiteTransforms {
		resultKite = append(resultKite, kite.applyTransformation(trans))
	}

	return resultDart, resultKite
}

var iterations = 0

const maxIterations = 4

func iterate(dartTranses []transformation, kiteTranses []transformation) ([]transformation, []transformation) {
	if iterations > maxIterations {
		return dartTranses, kiteTranses
	}
	iterations++

	newDartTranses, newKiteTranses := []transformation{}, []transformation{}

	for _, dartTrans := range dartTranses {
		tempDartTranses, tempKiteTranses := dartReplace(dartTrans)
		newDartTranses = append(newDartTranses, tempDartTranses...)
		newKiteTranses = append(newKiteTranses, tempKiteTranses...)
	}

	for _, kiteTrans := range kiteTranses {
		tempDartTranses, tempKiteTranses := kiteReplace(kiteTrans)
		newDartTranses = append(newDartTranses, tempDartTranses...)
		newKiteTranses = append(newKiteTranses, tempKiteTranses...)
	}

	return iterate(newDartTranses, newKiteTranses)
}

func kiteReplace(trans transformation) ([]transformation, []transformation) {
	kiteTranslate1 := coordinate{math.Cos(radian36), math.Sin(radian36)}
	kiteTranslate2 := coordinate{math.Cos(radian36), -math.Sin(radian36)}

	DartReplace := []transformation{
		combineTransform(trans, transformation{0, coordinate{0, 0}, 1}),
	}

	basicKiteReplace := []transformation{
		{6, kiteTranslate1.scale(trans.rescales).rotate(trans.amountOfRotation), 1},
		{-6, kiteTranslate2.scale(trans.rescales).rotate(trans.amountOfRotation), 1},
	}

	kiteReplaceTrans := []transformation{}
	for _, kiteTrans := range basicKiteReplace {
		kiteReplaceTrans = append(kiteReplaceTrans, combineTransform(trans, kiteTrans))
	}

	return DartReplace, kiteReplaceTrans
}

func dartReplace(trans transformation) ([]transformation, []transformation) {
	dartAngle := 3 * radian36

	dartTrans1 := coordinate{1 - math.Cos(dartAngle), -math.Sin(dartAngle)}
	dartTrans2 := coordinate{1 - math.Cos(-dartAngle), -math.Sin(-dartAngle)}

	basicDartTrans := []transformation{
		{3, dartTrans1.scale(trans.rescales + 1).rotate(trans.amountOfRotation), 1},
		{-3, dartTrans2.scale(trans.rescales + 1).rotate(trans.amountOfRotation), 1},
	}

	dartTranses := []transformation{}
	for _, dartTrans := range basicDartTrans {
		dartTranses = append(dartTranses, combineTransform(trans, dartTrans))
	}

	return dartTranses, []transformation{
		combineTransform(trans, transformation{1, coordinate{0, 0}, 1}),
		combineTransform(trans, transformation{-1, coordinate{0, 0}, 1}),
	}
}

