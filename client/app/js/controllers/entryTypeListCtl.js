'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('entryTypeListCtl', ['$scope', 'entryTypeService', '$state', function ($scope, entryTypeService, $state) {
        window.currentScope = $scope;
        var call = entryTypeService.List();
        call.then(function (data) {
            $scope.data = data;
        });
        $scope.edit = function (v) {
            $state.go('entryTypeEdit', { id: v });
        };
        $scope.create = function () {
            $state.go('entryTypeCreate');
        };
        $scope.formatter = function (value, row, index) {
            return [
                '<button type="submit" class="btn btn-primary" onclick="currentScope.edit(' + value + ')">Editar</button>',
            ].join('');
        };
    }]);
