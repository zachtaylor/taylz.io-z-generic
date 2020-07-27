package main

import _env "taylz.io/env"

func env() _env.Service {
	return _env.Service{
		"p": "{{P}}",
		"k": "{{TK}}",
		"v": "{{TV}}",
	}
}
