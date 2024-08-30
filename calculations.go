package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var startTime time.Time

const maxSecs = 5

type resultMutex struct {
	mu             sync.Mutex
	dartTransforms transSet
	kiteTransforms transSet
}

var result = resultMutex{dartTransforms: newSet(), kiteTransforms: newSet()}

func drawPolygons() ([]polygon, []polygon) {
	startTime = time.Now()
	dartTransforms, kiteTransforms := calculateDrawing()
	fmt.Println("Calculated!")
	resultDart, resultKite := []polygon{}, []polygon{}
	for trans := range dartTransforms.items {
		resultDart = append(resultDart, dart.applyTransformation(trans))
	}

	for trans := range kiteTransforms.items {
		resultKite = append(resultKite, kite.applyTransformation(trans))
	}

	return resultDart, resultKite
}

var iterations = 0

const maxIters = 10

func calculateDrawing() (transSet, transSet) {
	iterations = 0

	darts, kites := newSet(), newSet()

	for i := 0; i < 5; i++ {
		darts.add(transformation{2 * i, coordinate{0, 0}, 0})
	}

	for i := 0; i < maxIters; i++ {
		darts, kites = calculateNextStep(darts, kites)
	}

	return darts, kites
}

func calculateNextStep(dartTranses transSet, kiteTranses transSet) (transSet, transSet) {
	iterations++

	newDartTranses, newKiteTranses := newSet(), newSet()

	for dartTrans := range dartTranses.items {
		tempDartTranses, tempKiteTranses := dartReplace(dartTrans)
		newDartTranses = addAllNew(newDartTranses, tempDartTranses)
		newKiteTranses = addAllNew(newKiteTranses, tempKiteTranses)
	}

	for kiteTrans := range kiteTranses.items {
		tempDartTranses, tempKiteTranses := kiteReplace(kiteTrans)
		newDartTranses = addAllNew(newDartTranses, tempDartTranses)
		newKiteTranses = addAllNew(newKiteTranses, tempKiteTranses)
	}

	return newDartTranses, newKiteTranses
}

func addAllNew(transes transSet, toAdd []transformation) transSet {
	for _, temptrans := range toAdd {
		transes.add(temptrans)
	}
	return transes
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
