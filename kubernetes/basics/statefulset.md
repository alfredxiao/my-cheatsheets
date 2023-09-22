# Differences with Deployment
- Predictable and persistent Pod names (meaning each pod has its unique identity, and pods are not interchangable with each other)
- Predictable and persistent DNS hostnames
- Predictable and persistent volume bindings
- These three properties form the state of a Pod, sometimes referred to as its **`sticky ID`**.
- StatefulSets ensure this state/sticky-ID is persisted across failures, scaling, and other scheduling operations

# Notes
- `StatefulSet` name will become Pod dns name, e.g. `mystatefulset-0` `mystatefulset-1`
- **By default**, StatefulSets create one Pod at a time, and always **wait** for previous Pods to be running and ready before creating the next.
- Scaling operations are also governed by the same ordered startup rules.
  - For example, scaling from 3 to 5 replicas will start a new Pod called `mystatefulset-3` and wait for it the be running and ready before creating `mystatefulset-4`.
- Scaling down follows the **same rules in reverse** – the controller terminates the Pod with the highest index ordinal (number) first, waits for it to fully terminate before terminating the Pod with the next highest ordinal.
- StatefulSets currently require a **Headless Service** to be responsible for the network identity of the Pods. You are responsible for creating this Service.

# Deleting StatefulSet
- Firstly, deleting a `StatefulSet` does not terminate Pods in order.
- With this in mind, you may want to scale a StatefulSet to `0` replicas before deleting it.
- You can also use `terminationGracePeriodSeconds` to further control the way Pods are terminated.
  - It’s common to set this to at least 10 seconds to give applications running in Pods a chance to flush local buffers and safely commit any writes that are still “in-flight”.

# Volume
- any time a `StatefulSet` Pod fails or is terminated, associated volumes are unaffected. This allows replacement Pods to **attach to the same storage** as the Pods they’re replacing.
  - This is true, even if replacement Pods are scheduled to different cluster nodes.
  - This also means volume should be **remote** for `statefulset`, because local volume is attached to nodes.
- The same is true for scaling operations. If a `StatefulSet` Pod is deleted as part of a scale-down operation, subsequent scale-up operations will attach new Pods to the surviving volumes that match their names.
