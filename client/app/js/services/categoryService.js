'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.services')
    .service('categoryService', ['$http', '$q', 'config', function ($http, $q, config) {
        return {
            List: function () {
                var p = $http.get(config.serverUrl + '/category', { retries: config.retries });
                return defaultPromise(p, $q);
            },
            Get: function (id) {
                var p = $http.get(config.serverUrl + '/category/' + id, { retries: config.retries });
                return defaultPromise(p, $q);
            },
            Update: function (category) {
                var p = $http.post(config.serverUrl + '/category/update', category);
                return defaultPromise(p, $q);
            },
            Create: function (category) {
                var p = $http.post(config.serverUrl + '/category/create', category);
                return defaultPromise(p, $q);
            }
        };
    }]);
