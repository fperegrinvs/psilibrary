'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('adminCtl', ['$scope', 'Facebook', function ($scope, Facebook) {
 		call = Facebook.getUser(FB);
 		call.then(function(data){
 			$scope.user = data;
	     	console.log($scope.user);
 		});
    }]);
