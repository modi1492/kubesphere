/*

 Copyright 2019 The KubeSphere Authors.

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
package v1alpha2

import (
	"github.com/emicklei/go-restful"
	"github.com/emicklei/go-restful-openapi"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"kubesphere.io/kubesphere/pkg/apiserver/operations"
	"kubesphere.io/kubesphere/pkg/apiserver/runtime"
	"kubesphere.io/kubesphere/pkg/errors"
)

const GroupName = "operations.kubesphere.io"

var GroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}

var (
	WebServiceBuilder = runtime.NewContainerBuilder(addWebService)
	AddToContainer    = WebServiceBuilder.AddToContainer
)

func addWebService(c *restful.Container) error {

	tags := []string{"Operations"}

	webservice := runtime.NewWebService(GroupVersion)

	webservice.Route(webservice.POST("/nodes/{node}/drainage").
		To(operations.DrainNode).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("").
		Param(webservice.PathParameter("node", "node name")).
		Writes(errors.Error{}))

	webservice.Route(webservice.POST("/namespaces/{namespace}/jobs/{job}").
		To(operations.RerunJob).
		Metadata(restfulspec.KeyOpenAPITags, tags).
		Doc("Handle job operation").
		Param(webservice.PathParameter("job", "job name")).
		Param(webservice.PathParameter("namespace", "job's namespace")).
		Param(webservice.QueryParameter("a", "action")).
		Writes(""))

	c.Add(webservice)

	return nil
}
