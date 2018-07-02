# Oswee - Collaborative Resource Planning

GOlang backed Web desktop application. CRUD intensive
This is my personal learning project. My current main target is to learn best practices about structuring large scale project.
So I do research SCSS, Routing, Templating and other stuff first.

This repo is for my personal reference.

// 2018-03-25 Currently stopped any development and diving into Kubernetes world. Feeling really excited because this is the true way I want to manage my application.
Currently my target is to understand how to manage own Kubernates cluster, to deploy some applications in 3 modes - Dev, Stable and Edge. Then i want to understand how to run gRPC services on top of Vitess. I want fully separate my client from data layer and i would like to understand is it possible to isolate API from data source so i could change database later without heavy API redesigning.
Next thing will be something like static content CDN (JS, CSS, IMAGES).
And after that i need to learn how to do UI composition so i could develop stand alone modules and plug them as i need, to compose whole UI.

// 2018-03-31 Figured out some problem with persistent data volumes. Need to set up some server to server as data storage. Will try to look into some used SANs. Minikube is not solution for me and i got the idea, why it is not good idea to keep data into the same host. As temp solution will try to set up some VM.

// 2018-07-02 Started to think about API implementation. I have idea to start with ![gRPC icon] (https://grpc.io/img/grpc_square_reverse_4x.png) gRPC as main layer and then to add REST if required.