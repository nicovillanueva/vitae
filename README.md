# Vitae

## What
Here you can find the source code for my CV. The API is built in Go, and the .pdf, in LaTeX.

It is online at http://nicovillanueva.ddns.net/

The front end is under development, but the API is documented in [Swagger](http://nicovillanueva.ddns.net/swagger/index.html). View or download the pdf [here](http://nicovillanueva.ddns.net/api/download), or from the `latex` folder

## How
The API server is built using Echo, and the CLI using Cobra. It is packaged using Docker, and hosted in Google Cloud Engine. The DNS hostname is provided by NoIP.

Build locally running `make build`. This will generate the Swagger documentation, LaTeX, and Docker image. This requires [Swag](https://github.com/swaggo/swag) and [pdfLatex](https://www.tug.org/applications/pdftex/).

## Backlog
- Build process: Apply some kind of CI process. Travis CI, Docker automated builds, Gitlab CI
- Deployment: Deploying into Kubernetes would be overkill, but maybe to a better VM in AWS, GCE or DO. **Blocked for not having a job with which to pay hosting services!** :D
- HTTPS: Either embed a ssl-terminating proxy Nginx, or just Let's Encrypt
- Automatic DDNS updating: Just do it. Quick and easy
- API: Apply HATEOAS design. Not quick nor easy. Ugh.
- UI: Throw together an SPA (handmade, frameworkless).
- Automate LaTeX setup (container?)

## Contact
As mentioned in the API headers, contact me at `nicovillanueva|at|pm.me`, or via Linkedin, at https://www.linkedin.com/in/villanuevanicolas/
