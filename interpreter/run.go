package interpreter

// Run executes the environment.
func Run(env *Env) {
	for _, v := range env.SyntaxTree.ParseTree.Operations {
		RunTree(v, env)
	}
}
