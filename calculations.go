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

func (shape *polygon) applyTransformation(transform transformation) polygon {
	newPoints := []coordinate{}
	for _, coord := range shape.points {
		newPoint := applyTransformation(coord, transform)
		newPoints = append(newPoints, newPoint)
	}
	return polygon{newPoints}
}

func applyTransformation(coord coordinate, transform transformation) coordinate {
	// rescales
	rescale := math.Pow(scaleFactor, float64(transform.rescales))
	newCoord := coordinate{
		coord.x * rescale,
		coord.y * rescale,
	}

	// rotation
	angle := float64(transform.amountOfRotation) * radian36
	rotationMatrix := [][]float64{
		{math.Cos(angle), -math.Sin(angle)},
		{math.Sin(angle), math.Cos(angle)},
	}

	newCoord = matrixTransform(newCoord, rotationMatrix)

	// translation
	newCoord = coordinate{
		newCoord.x + transform.translation.x,
		newCoord.y + transform.translation.y,
	}

	return newCoord
}

func matrixTransform(coord coordinate, matrix [][]float64) coordinate {
	return coordinate{
		coord.x*matrix[0][0] + coord.y*matrix[0][1],
		coord.x*matrix[1][0] + coord.y*matrix[1][1],
	}
}

func kiteReplace() ([]transformation, []transformation) {
	return []transformation{
			{0, coordinate{0, 0}, 1},
		}, []transformation{
			{6, coordinate{math.Cos(radian36), math.Sin(radian36)}, 1},
			{-6, coordinate{math.Cos(radian36), -math.Sin(radian36)}, 1},
		}
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
