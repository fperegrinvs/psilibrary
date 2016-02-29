'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('entryTypeEditCtl', ['$scope', 'entryTypeService', '$state', function ($scope, entryTypeService, $state) {
        window.currentScope = $scope;
        var call = entryTypeService.Get($state.params.id);
        call.then(function (data) {
            $scope.data = data;
        });
        $scope.update = function () {
            entryTypeService.Update($scope.data).then(function () {
                $state.go('entryType');
            });
        };
    }]);
