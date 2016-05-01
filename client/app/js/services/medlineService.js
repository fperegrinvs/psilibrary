'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.services')
    .service('medlineService', ['$http', '$q', 'config', function ($http, $q, config) {
        return {
            Import: function (xml) {
                var p = $http.post(config.serverUrl + '/medline', xml);
                return defaultPromise(p, $q);
            },
        }
    }]);
