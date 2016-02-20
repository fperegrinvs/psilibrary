'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('categoryListCtl', ['$scope', 'categoryService', '$state', function($scope, categoryService, $state) {
        $scope.init = function() {
            window.currentScope = $scope;
            var call = categoryService.List();
            call.then(function(data) {
                $scope.data = data;
            });
        }
        $scope.edit = function(v) {
            $state.go('categoryEdit', { id: v });
        };
        $scope.create = function() {
            $state.go('categoryCreate');
        };
        $scope.formatter = function(value, row, index) {
            return [
                '<button type="submit" class="btn btn-primary" onclick="currentScope.edit(' + value + ')">Editar</button>',
            ].join('');
        };

        $scope.init();
    }]);
