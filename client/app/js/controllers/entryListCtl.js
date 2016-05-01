'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('entryListCtl', ['$scope', 'entryService', '$state', function ($scope, entryService, $state) {

    $scope.init = function() {
        window.currentScope = $scope;
        
        if ($state.params.error){
            $scope.msg = {error: $state.params.error}
        }

    	var call = entryService.List();
    	call.then(function(data){
    		$scope.data = data;
    	},
    	function(err){
    		$scope.msg = {error: err};
    	})
    }

    $scope.create = function() {
    	$state.go('entryCreate');
    }

    $scope.edit = function(id) {
    	if (id) {
	    	$state.go('entryEdit', { id: id });
    	}
    	else {
    		$scope.msg = {error: 'Nenhum registro selecionado'}
    	}
    }


    $scope.formatter = function (value, row, index) {
        return [
            '<button type="submit" class="btn btn-primary edit" onclick="currentScope.edit(' + value + ')">Editar</button>',
        ].join('');
    };

    $scope.abstractFormatter = function (value, row, index) {
        trimmedValue = value && value.length > 50 ? trimmedValue = value.substring(0, 50) + '...' : value;
        return trimmedValue;
    };

    $scope.init();
}]);
