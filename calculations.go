package main

import "math"

func drawPolygons() []polygon {
	dartTransforms, kiteTransforms := kiteReplace(transformation{0, coordinate{0, 0}, 0})

	newDartTransforms, newKiteTransforms := []transformation{}, []transformation{}

	for _, dartTrans := range dartTransforms {
		tempDartTranses, tempKiteTranses := dartReplace(dartTrans)
		newDartTransforms = append(newDartTransforms, tempDartTranses...)
		newKiteTransforms = append(newKiteTransforms, tempKiteTranses...)
	}

	for _, kiteTrans := range kiteTransforms {
		tempDartTranses, tempKiteTranses := kiteReplace(kiteTrans)
		newDartTransforms = append(newDartTransforms, tempDartTranses...)
		newKiteTransforms = append(newKiteTransforms, tempKiteTranses...)
	}

	result := []polygon{}
	for _, trans := range newDartTransforms {
		result = append(result, dart.applyTransformation(trans))
	}

	for _, trans := range newKiteTransforms {
		result = append(result, kite.applyTransformation(trans))
	}

	return result
}

func kiteReplace(trans transformation) ([]transformation, []transformation) {
	kiteTranslate1 := coordinate{math.Cos(radian36), math.Sin(radian36)}
	kiteTranslate2 := coordinate{math.Cos(radian36), -math.Sin(radian36)}

	DartReplace := []transformation{
		combineTransform(trans, transformation{0, coordinate{0, 0}, 1}),
	}

	basicKiteReplace := []transformation{
		{6, kiteTranslate1.scale(trans.rescales), 1},
		{-6, kiteTranslate2.scale(trans.rescales), 1},
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
		{3, dartTrans1.scale(trans.rescales + 1), 1},
		{-3, dartTrans2.scale(trans.rescales + 1), 1},
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

