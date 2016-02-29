'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('entryTypeCreateCtl', ['$scope', 'entryTypeService', '$state', function ($scope, entryTypeService, $state) {
        window.currentScope = $scope;
        $scope.data = {};
        $scope.update = function () {
            entryTypeService.Create($scope.data).then(function () {
                $state.go('entryType');
            });
        };
    }]);
