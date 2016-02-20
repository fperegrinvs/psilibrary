'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('categoryCreateCtl', ['$scope', 'categoryService', '$state', function($scope, categoryService, $state) {
		$scope.init = function() {
			window.currentScope = $scope;
			$scope.data = {};
		}
        $scope.update = function() {
            categoryService.Create($scope.data).then(function() {
                $state.go('category');
            });
        };

        $scope.init();
    }]);
