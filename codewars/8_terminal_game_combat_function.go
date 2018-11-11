package main

import "math"

// my

func combat(health, damage float64) float64 {
  if health - damage < 0 {
    return 0
  } else {
    return health - damage
  }
}

// best

func combat(health, damage float64) float64 {
  return math.Max(health-damage, 0.0)
}