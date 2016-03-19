'use strict'; // http://stackoverflow.com/questions/1335851/what-does-use-strict-do-in-javascript-and-what-is-the-reasoning-behind-it
angular.module('psilibrary.controllers')
    .controller('entryEditCtl', ['$scope', 'entryService', '$state', 'categoryService', 'entryTypeService', 
    	function ($scope, entryService, $state, categoryService, entryTypeService) {

	$scope.init = function(){
        if (!$state.params.id){
            $state.go('entryList', {error: 'Registro não encontrado'});
        }

        var entryCall = entryService.Get($state.params.id);
        entryCall.then(function(data){
            $scope.data = data;
        },
        function(err){
            $scope.msg = {error: err}
        });

        var entryCall = entryTypeService.List();
        entryCall.then(function(data){
            $scope.entryTypes = data;
        },
        function(err){
            $scope.msg = {error: err};
        });        

        var call = categoryService.List();
        call.then(function(data){
            $scope.categories = data;
        },
        function(err){
            $scope.msg = {error: err};
        });
	}

	$scope.save = function(){
		entryService.Update($scope.data).then(
			function(data){
				$scope.msg = {success: 'Registro atualizado com sucesso'};
			},
			function(err){
				$scope.msg = {error: err};
			});
	}

	$scope.init();
}]);