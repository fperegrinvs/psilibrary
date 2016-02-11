'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('categoryEditCtl', ['$scope', 'categoryService', '$state', function($scope, categoryService, $state) {
		window.currentScope = $scope;

		var call = categoryService.Get($state.params.id);
		call.then(function(data) {
			$scope.data = data;
		});

		$scope.update = function() {
			categoryService.Update($scope.data).then(function() {
				$state.go('category')
			});
		}
	}]);
