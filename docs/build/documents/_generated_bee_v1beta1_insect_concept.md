

-----------
# Bee v1beta1 insect

>bdocs-tab:example Bee Example

```bdocs-tab:example_yaml

apiVersion: insect.tamalsaha.com/v1beta1
kind: Bee
metadata:
  name: bee-example
spec:


```


Group        | Version     | Kind
------------ | ---------- | -----------
insect | v1beta1 | Bee







Bee

<aside class="notice">
Appears In:

<ul> 
<li><a href="#beelist-v1beta1-insect">BeeList insect/v1beta1</a></li>
</ul> </aside>

Field        | Description
------------ | -----------
apiVersion <br /> *string*    | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
kind <br /> *string*    | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
metadata <br /> *[ObjectMeta](#objectmeta-v1-meta)*    | 
spec <br /> *[BeeSpec](#beespec-v1beta1-insect)*    | 
status <br /> *[BeeStatus](#beestatus-v1beta1-insect)*    | 


### BeeSpec v1beta1 insect

<aside class="notice">
Appears In:

<ul>
<li><a href="#bee-v1beta1-insect">Bee insect/v1beta1</a></li>
</ul></aside>

Field        | Description
------------ | -----------

### BeeStatus v1beta1 insect

<aside class="notice">
Appears In:

<ul>
<li><a href="#bee-v1beta1-insect">Bee insect/v1beta1</a></li>
</ul></aside>

Field        | Description
------------ | -----------

### BeeList v1beta1 insect



Field        | Description
------------ | -----------
apiVersion <br /> *string*    | APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources
items <br /> *[Bee](#bee-v1beta1-insect) array*    | 
kind <br /> *string*    | Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds
metadata <br /> *[ListMeta](#listmeta-v1-meta)*    | 





