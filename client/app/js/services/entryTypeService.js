'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.services')
    .service('entryTypeService', ['$http', '$q', 'config', function ($http, $q, config) {
        return {
            List: function () {
                var p = $http.get(config.serverUrl + '/entrytype', { retries: config.retries });
                return defaultPromise(p, $q);
            },
            Get: function (id) {
                var p = $http.get(config.serverUrl + '/entrytype/' + id, { retries: config.retries });
                return defaultPromise(p, $q);
            },
            Update: function (entryType) {
                var p = $http.post(config.serverUrl + '/entrytype/update', entryType);
                return defaultPromise(p, $q);
            },
            Create: function (entryType) {
                var p = $http.post(config.serverUrl + '/entrytype/create', entryType);
                return defaultPromise(p, $q);
            }
        };
    }]);
