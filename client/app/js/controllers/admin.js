'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('adminCtl', ['$scope', 'userService', '$state', '$timeout', 'Facebook', '$rootScope'
      , function ($scope, userService, $state, $timeout, Facebook, $rootScope) {

      $scope.user = {};
      $scope.logged = false;
      $scope.byebye = false;
      $scope.salutation = false;

      // wait facebook to be ready      
      $scope.$watch(
        function() {
          return Facebook.isReady();
        },
        function(newVal) {
          if (newVal)
            $scope.facebookReady = true;
        }
      );
      
      var userIsConnected = false;
      
      Facebook.getLoginStatus(function(response) {
        if (response.status == 'connected') {
          userIsConnected = true;
          $scope.logged = true;
        }
      });
      
      // login if needed
      $scope.IntentLogin = function() {
        if(!userIsConnected) {
          $scope.login();
        }
      };
      

       // actual login
       $scope.login = function() {
         Facebook.login(function(response) {
          if (response.status == 'connected') {
            $scope.logged = true;
            $scope.me();
          }
        
        });
       };
       

       // user info
        $scope.me = function() {
          Facebook.api('/me', {fields: "id,name,picture"}, function(response) {
            /**
             * Using $scope.$apply since this happens outside angular framework.
             */
            $scope.$apply(function() {
              $scope.user = response;
              call = userService.Get(response.id);
              call.then(
                function(data){
                  if (data.id == response.login) {
                    $rootScope.authorized = true;
                    $rootScope.user = $scope.user;
                  }
                },
                function (data) {
                    $scope.clear();
                } 
              )
              console.log(response)
            });
            
          });
        };

      $scope.clear = function() {
        $rootScope.user = {};
        $scope.user = {}
        $rootScope.authorized = false;
      }
      
      // logout
      $scope.logout = function() {
        Facebook.logout(function() {
          $scope.$apply(function() {
            $scope.clear();
          });
        });
      }

      // check event change
      $scope.$on('Facebook:statusChange', function(ev, data) {
        if (data.status == 'connected') {
          $scope.$apply(function() {
            $scope.salutation = true;
            $scope.byebye     = false; 
            $scope.me();   
          });
        } else {
          $scope.$apply(function() {
            $scope.salutation = false;
            $scope.byebye     = true;
            
            // Dismiss byebye message after two seconds
            $timeout(function() {
              $scope.byebye = false;
            }, 2000)
          });
        }
      });
      }]);
