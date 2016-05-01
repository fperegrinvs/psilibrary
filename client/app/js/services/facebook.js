/*
angular.module('psilibrary.services')
.service('Facebook', ['$q', '$rootScope', function($q, $rootScope) {
  
  // resolving or rejecting a promise from a third-party
  // API such as Facebook must be
  // performed within $apply so that watchers get
  // notified of the change
  resolve = function(errval, retval, deferred) {
    $rootScope.$apply(function() {
      if (errval) {
        deferred.reject(errval);
      } else {
        retval.connected = true;
        deferred.resolve(retval);
      }
    });
  }
    
  return {
    getUser: function(FB) {
      var deferred = $q.defer();
      FB.getLoginStatus(function(response) {
        if (response.status == 'connected') {
          FB.api('/me', function(response) {
            resolve(null, response, deferred);
          });
        } else if (response.status == 'not_authorized') {
          FB.login(function(response) {
            if (response.authResponse) {
              FB.api('/me', function(response) {
                resolve(null, response, deferred);
              });
            } else {
              resolve(response.error, null, deferred);
            }
          });
        } 
      });
      promise = deferred.promise;
      promise.connected = false;
      return promise;
    }
  }; 
}]);
*/