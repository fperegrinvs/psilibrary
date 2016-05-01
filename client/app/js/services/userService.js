'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.services')
    .service('userService', ['$http', '$q', 'config', 'Facebook', '$state', '$timeout', function ($http, $q, config, Facebook, $state, $timeout) {
        return {
            Get: function (id) {
 				p = $http.get(config.serverUrl + '/user/' + id)
 				return defaultPromise(p, $q);
           }
        }
    }]);
