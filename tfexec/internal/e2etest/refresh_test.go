// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package e2etest

import (
	"context"
	"io"
	"regexp"
	"testing"

	"github.com/hashicorp/go-version"

	"github.com/aanogueira/terraform-exec/tfexec"
	"github.com/aanogueira/terraform-exec/tfexec/internal/testutil"
)

func TestRefresh(t *testing.T) {
	runTest(t, "basic", func(t *testing.T, tfv *version.Version, tf *tfexec.Terraform) {
		err := tf.Init(context.Background())
		if err != nil {
			t.Fatalf("error running Init in test directory: %s", err)
		}

		err = tf.Apply(context.Background())
		if err != nil {
			t.Fatalf("error running Apply: %s", err)
		}

		err = tf.Refresh(context.Background())
		if err != nil {
			t.Fatalf("error running Refresh: %s", err)
		}
	})
}

func TestRefreshJSON_TF014AndEarlier(t *testing.T) {
	versions := []string{testutil.Latest011, testutil.Latest012, testutil.Latest013, testutil.Latest014}

	runTestWithVersions(t, versions, "basic", func(t *testing.T, tfv *version.Version, tf *tfexec.Terraform) {
		err := tf.Init(context.Background())
		if err != nil {
			t.Fatalf("error running Init in test directory: %s", err)
		}

		re := regexp.MustCompile("terraform refresh -json was added in 0.15.3")

		err = tf.RefreshJSON(context.Background(), io.Discard)
		if err != nil && !re.MatchString(err.Error()) {
			t.Fatalf("error running Apply: %s", err)
		}
	})
}

func TestRefreshJSON_TF015AndLater(t *testing.T) {
	versions := []string{testutil.Latest015, testutil.Latest_v1, testutil.Latest_v1_1}

	runTestWithVersions(t, versions, "basic", func(t *testing.T, tfv *version.Version, tf *tfexec.Terraform) {
		err := tf.Init(context.Background())
		if err != nil {
			t.Fatalf("error running Init in test directory: %s", err)
		}

		err = tf.RefreshJSON(context.Background(), io.Discard)
		if err != nil {
			t.Fatalf("error running Apply: %s", err)
		}
	})
}
