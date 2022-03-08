package main

type contextKey string

const contextKeyIsAuthenticated = contextKey("isAuthenticated")
const contextKeyUser = contextKey("user")
