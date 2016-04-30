'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.services')
    .service('searchService', ['$http', '$q', 'config', function ($http, $q, config) {
        return {
            Search: function (query) {
                var p = $http.post(config.serverUrl + '/search', query);
                return defaultPromise(p, $q);
            },
        }
    }]);
