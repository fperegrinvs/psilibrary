'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it

angular.module('psilibrary.services')
    .service('entryTypeService', ['$http', '$q', 'config', function ($http, $q, config) {

	return {
		List : function() {
            var p = $http.get(config.serverUrl + '/entrytype', { retries: config.retries });

            var deferred = $q.defer();

            p.success(function (data, status) {
                deferred.resolve(data, status);
            });
            p.error(function (data, status) {
                deferred.reject(data, status);
            });

            return deferred.promise;
		},
		Get: function(id) {
            var p = $http.get(config.serverUrl + '/entrytype/' + id, { retries: config.retries });

            var deferred = $q.defer();

            p.success(function (data, status) {
                deferred.resolve(data, status);
            });
            p.error(function (data, status) {
                deferred.reject(data, status);
            });

            return deferred.promise;
		}
	};
 }]);
