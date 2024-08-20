package main

import "math"

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

func combineTransform(transFirst transformation, transSecond transformation) transformation {
	return transformation{
		transFirst.amountOfRotation + transSecond.amountOfRotation,
		coordinate{transFirst.translation.x + transSecond.translation.x, transFirst.translation.y + transSecond.translation.y},
		transFirst.rescales + transSecond.rescales,
	}
}

func (coord *coordinate) scale(scaleExp int) coordinate {
	return coordinate{
		coord.x * math.Pow(scaleFactor, float64(scaleExp)),
		coord.y * math.Pow(scaleFactor, float64(scaleExp)),
	}
}
