'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.services')
    .service('entryService', ['$http', '$q', 'config', function ($http, $q, config) {
        return {
            List: function () {
                var p = $http.get(config.serverUrl + '/entry', { retries: config.retries });
                return defaultPromise(p, $q);
            },
            Get: function (id) {
                var p = $http.get(config.serverUrl + '/entry/' + id, { retries: config.retries });
                return defaultPromise(p, $q);
            },
            Update: function (entry) {
                var p = $http.post(config.serverUrl + '/entry/update', entry);
                return defaultPromise(p, $q);
            },
            Create: function (entry) {
                var p = $http.post(config.serverUrl + '/entry/create', entry);
                return defaultPromise(p, $q);
            }
        }
    }]);
