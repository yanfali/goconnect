#GoConnect Web Application
##Directory Layout and Organization
This is going to be opinionated. The top level of this directory is the basis of the web application root directory.

* /root - of application
	* ```bin``` - where the web application binary lives, you run this
	* ```src```
		* ```webapp```
			* ```config``` - project specific config object
			* ```handlers``` - ServeHTTP go code. Each screen should have at least 1 go handler. There may be more than 1 go handler per page. These all belong to the package handlers. By convention, handlers should have a camel case name of {{Action}}Handler. e.g. LoginHandler
			* main.go - configuration of the web application through code for now.
	* ```webapp``` - where the web application assets live
		* ```images```
			* ```common``` - shared images
			* and optionally one directory per Handler type name, e.g. ```LoginHandler```
		* scripts
			* ```common``` - shared javascript assets, preferrably by lib name, e.g. ```common/requirejs```
			* and optionally one directory per Handler type name, e.g. ```LoginHandler```
		* styles
			* ```common``` - shared style sheets
			* and optionally one directory per Handler type name, e.g. ```LoginHandler``` with handler specific style sheets
		* templates - golang html/templates
			* common - things like html5 boiler plate for the site, shared forms, etc...
			* and optionally one directory per Handler type name, e.g. ```LoginHandler``` with handler specific style sheets. The default template is called ```body.tmpl```. You use this template to bootstrap all your other templates. For now it lives inside a div with id of ```app```.

##TODO

* add hooks to head so you can add custom content or templates
* possibly make the handler own ```body``` instead of a ```div``` inside of ```body``` called ```app```
* create a default LoginHandler which can easily be replaced, or customized for the specific project