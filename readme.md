# ScalerD
ScalerD also called "scaler daemon" is a service that spawns local program instances until some CPU utilization threshold is reached.

This allows a server to maximize the number of client services (typically microservice instances) that are running to take advantage of as much vertical scaling capacity as possible.

When ScalerD starts, it checks its configuration for a list of services that it will be managing.
An instance of each service is started, and the instance's PID is registered within ScalerD.

ScalerD will continue to spawn new service instances at a configured rate until the CPU threshold is reached.