//
// Last.Backend LLC CONFIDENTIAL
// __________________
//
// [2014] - [2017] Last.Backend LLC
// All Rights Reserved.
//
// NOTICE:  All information contained herein is, and remains
// the property of Last.Backend LLC and its suppliers,
// if any.  The intellectual and technical concepts contained
// herein are proprietary to Last.Backend LLC
// and its suppliers and may be covered by Russian Federation and Foreign Patents,
// patents in process, and are protected by trade secret or copyright law.
// Dissemination of this information or reproduction of this material
// is strictly forbidden unless prior written permission is obtained
// from Last.Backend LLC.
//

package service_test

//
//import (
//	"github.com/lastbackend/lastbackend/pkg/cli/cmd/service"
//	"github.com/lastbackend/lastbackend/pkg/cli/context"
//	h "github.com/lastbackend/lastbackend/pkg/util/http"
//	"github.com/stretchr/testify/assert"
//	"net/http"
//	"net/http/httptest"
//	"testing"
//)
//
//func TestLogs(t *testing.T) {
//
//	const (
//		app   string = "app"
//		serviceName   string = "service"
//		podName       string = "pod"
//		containerName string = "container"
//		token         string = "mocktoken"
//	)
//
//	var (
//		err error
//		ctx = context.Mock()
//	)
//
//	ctx.Token = token
//
//	//------------------------------------------------------------------------------------------
//	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//
//		tk := r.Header.Get("Authorization")
//		assert.NotEmpty(t, tk, "token should be not empty")
//		assert.Equal(t, tk, "Bearer "+token, "they should be equal")
//
//		var (
//			query          = r.URL.Query()
//			podQuery       = query.Get("pod")
//			containerQuery = query.Get("container")
//		)
//
//		assert.Equal(t, podQuery, podName, "they should be equal")
//		assert.Equal(t, containerQuery, containerName, "they should be equal")
//
//		w.WriteHeader(200)
//	}))
//	defer server.Close()
//	//------------------------------------------------------------------------------------------
//
//	ctx.HTTP = h.New(server.URL)
//
//	_, err = service.Logs(app, serviceName, podName, containerName)
//	if err != nil {
//		t.Error(err)
//	}
//}
