exports.config =
  # See http://brunch.io/#documentation for docs.
  modules:
    wrapper: false
    definition: false
  files:
    javascripts:
      joinTo: 'app.js'
    stylesheets: {
        joinTo: {
            'css/vendor.css': /^(vendor|bower_components)\//,
            'css/styles.css': /^app\/css\//
        }
    }
    templates:
      joinTo: 'templates.js'
  plugins:{
    copycat:{
      "fonts" : ["bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.eot", "bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.svg", "bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.ttf", "bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.woff", "bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.woff2" ],
      verbose : true,
      onlyChanged: true
    }
  }
