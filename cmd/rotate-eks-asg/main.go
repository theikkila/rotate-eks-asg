package main

import (
	"log"
	"os"

	"github.com/complex64/go-utils/pkg/ctxutil"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/tenjin/rotate-eks-asg/internal/pkg/rotator"
)

var (
	cluster = kingpin.Arg("asg", "EKS cluster formed by the Auto Scaling Groups (ASG)").Required().String()
	groups  = kingpin.Arg("cluster", "EKS Auto Scaling Groups to rotate").Required().Strings()
)

func init() { // TODO
	_ = os.Setenv("AWS_SDK_LOAD_CONFIG", "true")
}

func main() {
	kingpin.Parse()
	ctx, cancel := ctxutil.ContextWithCancelSignals(os.Kill, os.Interrupt)
	defer cancel()
	if err := rotator.RotateAll(ctx, *cluster, *groups); err != nil {
		log.Fatal(err)
	}
}
