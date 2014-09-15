package main

import (
	"fmt"
)

func main() {
	q := 96
	src := []string{
		`package main`,
		``,
		`import (`,
		`    "fmt"`,
		`)`,
		``,
		`func main(){`,
		`    q := 96`,
		`    src := []string {`,
		`        `,
		`    }`,
		`    for i := 0; i < 9; i++ {`,
		`        fmt.Println(src[i])`,
		`    }`,
		`    for i := 0; i < len(src); i++ {`,
		`        fmt.Println(src[9] + string(q) + src[i] + string(q) + ",")`,
		`    }`,
		`    for i := 10; i < len(src); i++ {`,
		`        fmt.Println(src[i])`,
		`    }`,
		`}`,
		``,
	}
	for i := 0; i < 9; i++ {
		fmt.Println(src[i])
	}
	for i := 0; i < len(src); i++ {
		fmt.Println(src[9] + string(q) + src[i] + string(q) + ",")
	}
	for i := 10; i < len(src); i++ {
		fmt.Println(src[i])
	}
}
