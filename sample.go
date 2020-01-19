package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	maxNumOfTrials = 1000
)

type Agent struct {
	Trials int
	Reward [][]int
}

type Arm struct {
	Prob  float64
	Count int
}

type Arms []Arm

func test() {
	fmt.Println("UCB1 algorithm")

	p := []float64{0.28, 0.4, 0.6, 0.15}
	lineCount := len(p)

	agent := Agent{}
	agent.Trials = 0

	agent.Reward = make([][]int, lineCount)
	for i := range agent.Reward {
		agent.Reward[i] = make([]int, maxNumOfTrials)
	}
	for i := 0; i < lineCount; i++ {
		for j := 0; j < maxNumOfTrials; j++ {
			agent.Reward[i][j] = 0
		}
	}
	arm := make(Arms, lineCount)
	for i := 0; i < lineCount; i++ {
		arm[i].Prob = p[i]
		arm[i].Count = 0
	}

	for ; agent.Trials < maxNumOfTrials; agent.Trials++ {
		s := UCB1(agent, arm, lineCount)

		reward := Bernoulli_try(&arm[s])
		//agent.Reward[s][arm[s].Count - 1] = reward
		agent.Reward[s][agent.Trials] = reward

		PrintStats(agent, arm, s, reward, lineCount)

	}
}

func UCB1(agent Agent, arm Arms, count int) int {
	for i := 0; i < count; i++ {
		if arm[i].Count == 0 {
			return i
		}
	}

	x_mean := Sample_Mean(agent, arm, count)

	//fmt.Println(x_mean)
	//choose the best arm index
	val := make([]float64, count)
	s := 0
	max := 0.0
	for i := 0; i < count; i++ {
		val[i] = x_mean[i] + math.Sqrt(math.Log(float64(agent.Trials))/(2*float64(arm[i].Count)))
		if x_mean[i] != 0 {
			fmt.Println("olalal")
		}
		if val[i] > max {
			max = val[i]
			s = i
		}
	}
	return s
}

func Bernoulli_try(arm *Arm) int {
	if arm.Prob < 0 || arm.Prob > 1 {
		return -1
	}

	rand.Seed(time.Now().UnixNano())

	d := rand.Float64()

	arm.Count++

	if arm.Prob > d {
		return 1
	} else {
		return 0
	}

}

func Sample_Mean(agent Agent, arm Arms, count int) []float64 {
	x_mean := make([]float64, count)
	for i := 0; i < count; i++ {
		if arm[i].Count != 0 {
			for j := range agent.Reward {
				x_mean[i] = x_mean[i] + float64(agent.Reward[i][j])
			}
			x_mean[i] = x_mean[i] / float64(arm[i].Count)
		} else {
			x_mean[i] = 0.0
		}
	}
	return x_mean
}

func PrintStats(agent Agent, arm Arms, s int, reward int, lineCount int) {
	fmt.Println("------------------------------")
	fmt.Println("Total Trials: ", agent.Trials+1)
	fmt.Println("Selected arm:", s)
	fmt.Println("Reward:", reward)
	fmt.Println("[{Success probability, Number of selected}] = ", arm)
	fmt.Println("[Count]")
	for i := 0; i < lineCount; i++ {
		fmt.Println("Arm:", i, "=>", arm[i].Count, "times")
	}
	fmt.Println("------------------------------")
	for i := 0; i < lineCount; i++ {
		fmt.Println("Arm:", i, "=>", arm[i].Prob, "prob")
	}
	SumReward := make([]float64, lineCount)
	for i := 0; i < agent.Trials+1; i++ {
		for j := 0; j < lineCount; j++ {
			SumReward[j] = SumReward[j] + float64(agent.Reward[j][i])
		}
	}
	for i := 0; i < lineCount; i++ {
		SumReward[i] = SumReward[i] / float64(arm[i].Count)

	}
	fmt.Println("Total Rewards:", SumReward)
	fmt.Println("------------------------------")

}
