/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package metrics

import (
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/require"
)

func TestMetrics(t *testing.T) {
	m := Get()
	require.NotNil(t, m)
	require.True(t, m == Get())

	t.Run("ActivityPub", func(t *testing.T) {
		require.NotPanics(t, func() { m.InboxHandlerTime("Create", time.Second) })
		require.NotPanics(t, func() { m.OutboxPostTime(time.Second) })
		require.NotPanics(t, func() { m.OutboxResolveInboxesTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorBuildCredentialTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorSignCredentialTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorPostOfferActivityTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorGetWitnessesTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorGetPreviousAnchorsTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorGetPreviousAnchorsGetBulkTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorSignWithLocalWitnessTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorSignWithServerKeyTime(time.Second) })
		require.NotPanics(t, func() { m.WitnessAnchorCredentialTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorSignLocalWitnessLogTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorSignLocalStoreTime(time.Second) })
		require.NotPanics(t, func() { m.WriteAnchorSignLocalWatchTime(time.Second) })
		require.NotPanics(t, func() { m.ProcessWitnessedAnchorCredentialTime(time.Second) })
		require.NotPanics(t, func() { m.AddOperationTime(time.Second) })
		require.NotPanics(t, func() { m.BatchCutTime(time.Second) })
		require.NotPanics(t, func() { m.BatchRollbackTime(time.Second) })
		require.NotPanics(t, func() { m.BatchAckTime(time.Second) })
		require.NotPanics(t, func() { m.BatchNackTime(time.Second) })
		require.NotPanics(t, func() { m.BatchSize(float64(500)) })
		require.NotPanics(t, func() { m.ProcessAnchorTime(time.Second) })
		require.NotPanics(t, func() { m.ProcessDIDTime(time.Second) })
		require.NotPanics(t, func() { m.CASWriteTime(time.Second) })
		require.NotPanics(t, func() { m.CASResolveTime(time.Second) })
		require.NotPanics(t, func() { m.CASIncrementCacheHitCount() })
		require.NotPanics(t, func() { m.CASReadTime("local", time.Second) })
		require.NotPanics(t, func() { m.DocumentCreateUpdateTime(time.Second) })
		require.NotPanics(t, func() { m.DocumentResolveTime(time.Second) })
		require.NotPanics(t, func() { m.OutboxIncrementActivityCount("Create") })
		require.NotPanics(t, func() { m.DBPutTime("CouchDB", time.Second) })
		require.NotPanics(t, func() { m.DBGetTime("CouchDB", time.Second) })
		require.NotPanics(t, func() { m.DBGetTagsTime("CouchDB", time.Second) })
		require.NotPanics(t, func() { m.DBGetBulkTime("CouchDB", time.Second) })
		require.NotPanics(t, func() { m.DBQueryTime("CouchDB", time.Second) })
		require.NotPanics(t, func() { m.DBDeleteTime("CouchDB", time.Second) })
		require.NotPanics(t, func() { m.WitnessAddProofVctNil(time.Second) })
		require.NotPanics(t, func() { m.WitnessAddVC(time.Second) })
		require.NotPanics(t, func() { m.WitnessAddProof(time.Second) })
		require.NotPanics(t, func() { m.WitnessWebFinger(time.Second) })
		require.NotPanics(t, func() { m.WitnessVerifyVCTSignature(time.Second) })
		require.NotPanics(t, func() { m.AddProofParseCredential(time.Second) })
		require.NotPanics(t, func() { m.AddProofSign(time.Second) })
		require.NotPanics(t, func() { m.SignerGetKey(time.Second) })
		require.NotPanics(t, func() { m.SignerSign(time.Second) })
		require.NotPanics(t, func() { m.SignerAddLinkedDataProof(time.Second) })
	})
}

func TestNewCounter(t *testing.T) {
	labels := prometheus.Labels{"type": "create"}

	require.NotNil(t, newCounter("activityPub", "metric_name", "Some help", labels))
}

func TestNewHistogram(t *testing.T) {
	labels := prometheus.Labels{"type": "create"}

	require.NotNil(t, newHistogram("activityPub", "metric_name", "Some help", labels))
}

func TestNewGuage(t *testing.T) {
	require.NotNil(t, newGauge("activityPub", "metric_name", "Some help"))
}
