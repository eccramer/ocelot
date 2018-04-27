package builder

import (
	"bitbucket.org/level11consulting/ocelot/build"
	"bitbucket.org/level11consulting/ocelot/models/pb"
	"golang.org/x/net/context"

	"testing"
)


// test that in docker, can run the InstallPackageDeps to multiple image types
func TestDockerBasher_InstallPackageDeps_alpine36(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping due to -short flag being set")
	}
	ctx := context.Background()
	alpine, cleanupFunc := CreateLivingDockerContainer(t, ctx, "alpine:3.6")
	defer cleanupFunc(t)
	su := build.InitStageUtil("alpine36test")
	logout := make(chan[]byte, 10000)
	result := alpine.Exec(ctx, su.GetStage(), su.GetStageLabel(), []string{}, alpine.InstallPackageDeps(), logout)
	t.Log(result.Status)
	t.Log(string(<-logout))
	if result.Status == pb.StageResultVal_FAIL {
		t.Error("couldn't download deps! oh nuuu!")
		return
	}
	testDeps := []string{"/bin/sh", "-c", "command -v openssl && command -v bash && command -v zip && command -v wget && command -v python"}
	result = alpine.Exec(ctx, su.GetStage(), su.GetStageLabel(), []string{}, testDeps, logout)
	if result.Status == pb.StageResultVal_FAIL {
		t.Error("deps not found! oh nuuu!")
	}
	t.Log(result.Status)
	t.Log(string(<-logout))
}
//
//// test that in docker, can run the InstallPackageDeps to multiple image types
func TestDockerBasher_InstallPackageDeps_alpinelatest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping due to -short flag being set")
	}
	ctx := context.Background()
	alpine, cleanupFunc := CreateLivingDockerContainer(t, ctx, "alpine:latest")
	defer cleanupFunc(t)
	su := build.InitStageUtil("alpinelatestTest")
	logout := make(chan[]byte, 10000)
	result := alpine.Exec(ctx, su.GetStage(), su.GetStageLabel(), []string{}, alpine.InstallPackageDeps(), logout)
	t.Log(result.Status)
	t.Log(string(<-logout))
	if result.Status == pb.StageResultVal_FAIL {
		t.Error("couldn't download deps! oh nuuu!")
		return
	}
	testDeps := []string{"/bin/sh", "-c", "command -v openssl && command -v bash && command -v zip && command -v wget && command -v python"}
	result = alpine.Exec(ctx, su.GetStage(), su.GetStageLabel(), []string{}, testDeps, logout)
	if result.Status == pb.StageResultVal_FAIL {
		t.Error("deps not found! oh nuuu!")
	}
	t.Log(result.Status)
	t.Log(string(<-logout))
}


// test that in docker, can run the InstallPackageDeps to multiple image types
func TestDockerBasher_InstallPackageDeps_ubuntu1604(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping due to -short flag being set")
	}
	ctx := context.Background()
	alpine, cleanupFunc := CreateLivingDockerContainer(t, ctx, "ubuntu:16.04")
	defer cleanupFunc(t)
	su := build.InitStageUtil("ubuntu1604test")
	logout := make(chan[]byte, 10000)
	result := alpine.Exec(ctx, su.GetStage(), su.GetStageLabel(), []string{}, alpine.InstallPackageDeps(), logout)
	t.Log(result.Status)
	t.Log(string(<-logout))
	if result.Status == pb.StageResultVal_FAIL {
		t.Error("couldn't download deps! oh nuuu!")
		return
	}
	testDeps := []string{"/bin/sh", "-c", "command -v openssl && command -v bash && command -v zip && command -v wget && command -v python"}
	result = alpine.Exec(ctx, su.GetStage(), su.GetStageLabel(), []string{}, testDeps, logout)
	if result.Status == pb.StageResultVal_FAIL {
		t.Error("deps not found! oh nuuu!")
	}
	t.Log(result.Status)
	t.Log(string(<-logout))
}

