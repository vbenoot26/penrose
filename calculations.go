package main

import (
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

var result = resultMutex{dartTransforms: make(transSet), kiteTransforms: make(transSet)}

var iterations = 0

const maxIters = 10

func calculateDrawings() []state {
	iterations = 0

	darts, kites := initSun()

	result := []state{{darts, kites}}
	for i := 0; i < maxIters; i++ {
		newDarts, newKites := calculateNextStep(result[i].dartTranses, result[i].kiteTranses)
		result = append(result, state{newDarts, newKites})
		darts, kites = calculateNextStep(darts, kites)
	}

	return result
}

func initDart() (transSet, transSet) {
	return *createSet(idTransform), make(transSet)
}

func initKite() (transSet, transSet) {
	return make(transSet), *createSet(idTransform)
}

func initSun() (transSet, transSet) {
	darts, kites := make(transSet), make(transSet)

	for i := 0; i < 5; i++ {
		darts.add(transformation{2 * i, coordinate{0, 0}, 0})
	}

	return darts, kites
}

func calculateNextStep(dartTranses transSet, kiteTranses transSet) (transSet, transSet) {
	iterations++

	newDartTranses, newKiteTranses := make(transSet), make(transSet)

	for dartTrans := range dartTranses {
		tempDartTranses, tempKiteTranses := dartReplace(dartTrans)
		newDartTranses = addAllNew(newDartTranses, tempDartTranses)
		newKiteTranses = addAllNew(newKiteTranses, tempKiteTranses)
	}

	for kiteTrans := range kiteTranses {
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
		{6, kiteTranslate1.scaleByFactor(trans.rescales).rotate(trans.amountOfRotation), 1},
		{-6, kiteTranslate2.scaleByFactor(trans.rescales).rotate(trans.amountOfRotation), 1},
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
		{3, dartTrans1.scaleByFactor(trans.rescales + 1).rotate(trans.amountOfRotation), 1},
		{-3, dartTrans2.scaleByFactor(trans.rescales + 1).rotate(trans.amountOfRotation), 1},
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
