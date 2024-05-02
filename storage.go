package main

type Storage interface {
  Get(int) (any, error)
  Put()
}
