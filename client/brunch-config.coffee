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
      "fonts" : ["bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.eot", "bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.svg", "bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.ttf", "bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.woff", "bower_components/bootstrap/dist/fonts/glyphicons-halflings-regular.woff2",  "bower_components/font-awesome/fonts/fontawesome-webfont.eot",  "bower_components/font-awesome/fonts/fontawesome-webfont.svg",  "bower_components/font-awesome/fonts/fontawesome-webfont.ttf",  "bower_components/font-awesome/fonts/fontawesome-webfont.woff",  "bower_components/font-awesome/fonts/FontAwesome.otf"],
      verbose : true,
      onlyChanged: true
    }
  }
