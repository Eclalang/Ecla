package interpreter

// Run executes the environment.
func Run(env *Env) {
	for _, v := range env.SyntaxTree.ParseTree.Operations {
		RunTree(v, env)
	}
}

// Load executes import statements with their environment.
func Load(env *Env) {
	for _, v := range env.SyntaxTree.ParseTree.Operations {
		RunTreeLoad(v, env)
	}
}
