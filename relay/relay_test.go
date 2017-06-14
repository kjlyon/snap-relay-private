/*
http://www.apache.org/licenses/LICENSE-2.0.txt

Copyright 2017 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package relay

import (
	"testing"
	"time"

	"github.com/intelsdi-x/snap/control/plugin/client"
	. "github.com/smartystreets/goconvey/convey"
)

func TestRelay(t *testing.T) {

	_, err := client.NewStreamCollectorGrpcClient(
		"localhost:8181",
		5*time.Second,
		client.SecurityTLSOff(),
	)
	if err != nil {
		panic(err)
	}
	Convey("Test StreamMetrics", t, func() {

		// cfg := cdata.NewNode()
		// cfg.AddItem("MaxCollectDuration", ctypes.ConfigValueInt{Value: 5000000000})
		// cfg.AddItem("MaxMetricsBuffer", ctypes.ConfigValueInt{Value: 2})
		// requested_metrics := []core.Metric{
		// 	plugin.MetricType{
		// 		Namespace_: core.NewNamespace("animal", "cat"),
		// 		Config_:    cfg,
		// 	},
		// }

		//rq <- requested_metrics

		// metricsOut, errOut, err := c.StreamMetrics(requested_metrics)
		// So(metricsOut, ShouldNotBeEmpty)
		// So(errOut, ShouldBeNil)
		// So(err, ShouldBeNil)

	})

	// Convey("Test GetMetricTypes", t, func() {
	// 	r := relay{}

	// 	Convey("Collect String", func() {
	// 		mt, err := r.GetMetricTypes(nil)
	// 		So(err, ShouldBeNil)
	// 		So(len(mt), ShouldEqual, 2)
	// 	})

	// })

	// Convey("Test GetConfigPolicy", t, func() {
	// 	r := relay{}
	// 	_, err := r.GetConfigPolicy()

	// 	Convey("No error returned", func() {
	// 		So(err, ShouldBeNil)
	// 	})

	// })

}
