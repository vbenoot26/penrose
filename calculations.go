package main

import (
	"context"
	"fmt"
	"math"
	"sync"
	"time"
)

var startTime time.Time

const maxSecs = 500 * time.Millisecond

type resultMutex struct {
	mu             sync.Mutex
	dartTransforms []transformation
	kiteTransforms []transformation
}

var result = resultMutex{dartTransforms: []transformation{}, kiteTransforms: []transformation{}}

func (result *resultMutex) setResults(dartResults []transformation, kiteResults []transformation) {
	result.mu.Lock()
	defer result.mu.Unlock()

	result.dartTransforms = dartResults
	result.kiteTransforms = kiteResults
}

func (result *resultMutex) getResults() ([]transformation, []transformation) {
	result.mu.Lock()
	defer result.mu.Unlock()
	return result.dartTransforms, result.kiteTransforms
}

func drawPolygons() ([]polygon, []polygon) {
	startTime = time.Now()
	dartTransforms, kiteTransforms := calculateDrawing()
	fmt.Println("Calculated!")
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

func calculateDrawing() ([]transformation, []transformation) {
	iterations = 0

	ctx, cancel := context.WithTimeout(context.Background(), maxSecs*time.Second)
	defer cancel()

	go func() {
		iterate([]transformation{{0, coordinate{0, 0}, 0}}, []transformation{}, ctx)
	}()

	time.Sleep(maxSecs * time.Second)
	fmt.Println(iterations)
	return result.getResults()
}

func iterate(dartTranses []transformation, kiteTranses []transformation, ctx context.Context) {
	iterations++
	fmt.Println(iterations)

	newDartTranses, newKiteTranses := []transformation{}, []transformation{}

	for _, dartTrans := range dartTranses {
		tempDartTranses, tempKiteTranses := dartReplace(dartTrans)
		newDartTranses = addAllNew(newDartTranses, tempDartTranses)
		newKiteTranses = addAllNew(newKiteTranses, tempKiteTranses)
	}

	for _, kiteTrans := range kiteTranses {
		tempDartTranses, tempKiteTranses := kiteReplace(kiteTrans)
		newDartTranses = addAllNew(newDartTranses, tempDartTranses)
		newKiteTranses = addAllNew(newKiteTranses, tempKiteTranses)
	}

	select {
	case <-ctx.Done():
		return

	default:
		result.setResults(newDartTranses, newKiteTranses)
		iterate(newDartTranses, newKiteTranses, ctx)
	}
}

func addAllNew(transes []transformation, toAdd []transformation) []transformation {
	resultslice := []transformation{}
	copy(transes, resultslice)
	for _, temptrans := range toAdd {
		resultslice = addIfNotContains(resultslice, temptrans)
	}
	return resultslice
}

func addIfNotContains(transes []transformation, trans transformation) []transformation {
	resultSlice := []transformation{}
	copy(resultSlice, transes)
	for _, tempTrans := range transes {
		if trans == tempTrans {
			return resultSlice
		}
	}
	return append(resultSlice, trans)
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
