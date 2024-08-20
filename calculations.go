package main

import "math"

func drawPolygons() []polygon {
	dartTransforms, kiteTransforms := dartReplace()

	result := []polygon{}
	for _, trans := range dartTransforms {
		result = append(result, dart.applyTransformation(trans))
	}

	for _, trans := range kiteTransforms {
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

func dartReplace() ([]transformation, []transformation) {
	dartAngle := 3 * radian36

	return []transformation{
			{3, coordinate{scaleFactor * (1 - math.Cos(dartAngle)), scaleFactor * -math.Sin(dartAngle)}, 1},
			{-3, coordinate{scaleFactor * (1 - math.Cos(-dartAngle)), scaleFactor * -math.Sin(-dartAngle)}, 1},
		}, []transformation{
			{1, coordinate{0, 0}, 1},
			{-1, coordinate{0, 0}, 1},
		}
}
