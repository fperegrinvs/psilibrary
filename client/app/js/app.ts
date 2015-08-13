
/*================================================================
=>                  App = client
==================================================================*/
/*global angular*/

var app = angular.module('client', ["ngCookies", "ngResource", "ngSanitize", "ngRoute", "ngAnimate", "ui.utils", "ui.bootstrap", "ui.router", "psilibrary.controllers", 'psilibrary.templates', 'psilibrary.config', 'psilibrary.services', 'psilibrary.directives']);

angular.module("psilibrary.controllers", []);
angular.module("psilibrary.templates", []);
angular.module("psilibrary.services", []);
angular.module("psilibrary.config", []);
angular.module("psilibrary.directives", []);



app.config(['$routeProvider', '$locationProvider', '$httpProvider', '$stateProvider','$urlRouterProvider', function($routeProvider, $locationProvider, $httpProvider, $stateProvider, $urlRouterProvider) {
	'use strict';
    //Enable cross domain calls

    $urlRouterProvider.otherwise('/404');


    var createState = function(name, url, partial, controllerName, view) {
        var views = {};
        views[view + "@"] = { templateUrl: "templates/" + partial + ".html", controller: controllerName + "Ctl" };
        return $stateProvider.state(name, { url: url, views: views });
    };

    var stateTree = [
        'home',
        'home2',
        'entryType',
        '404',
    ];

    var stateConfigs = {
        'home': { url: '', partial: 'home', controller: 'home', target: 'miolo' },
        'home2': { url: '/', partial: 'home', controller: 'home', target: 'miolo' },
        'entryType': { url: '/entryType', controller: 'entryTypeList', partial: 'entryTypeList', target: 'miolo'},
        '404': { url: '/404', partial: '404', controller: 'home', target: 'miolo' },
    };

    stateTree = addSubStates(stateTree);
    var stateoptions = createStateOptions(stateTree, stateConfigs);

    stateoptions.forEach(function(state) { createState(state.name, state.url, state.partial, state.controller, state.target); });

	// This is required for Browser Sync to work poperly
	$httpProvider.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
}])

	.config(['$locationProvider', function($location) {
		$location.hashPrefix('!');
	}])

	.run(['$rootScope', '$state', '$stateParams', function($rootScope, $state, $stateParams) {
		// It's very handy to add references to $state and $stateParams to the $rootScope
		// so that you can access them from any scope within your applications.For example,
		// <li ui-sref-active="active }"> will set the <li> // to active whenever
		// 'contacts.list' or one of its decendents is active.
		$rootScope.$state = $state;
		$rootScope.$stateParams = $stateParams;
		$rootScope.path = window.path;

		$rootScope.$on('$stateChangeStart', function(event, toState, toParams, fromState, fromParams) {
			console.log('$stateChangeStart to ' + toState.to + '- fired when the transition begins. toState,toParams : \n', toState, toParams);
		});
		$rootScope.$on('$stateChangeError', function(event, toState, toParams, fromState, fromParams) {
			console.log('$stateChangeError - fired when an error occurs during transition.');
			console.log(arguments);
		});
		$rootScope.$on('$stateChangeSuccess', function(event, toState, toParams, fromState, fromParams) {
			console.log('$stateChangeSuccess to ' + toState.name + '- fired once the state transition is complete.');
		});
	}]);



function addSubStates(stateTree) {
    var entries = {};

    // mapeando dependencias
    stateTree.forEach(function (name) {
        if (name.split('.')[0] == '*') {
            var master = name.split('.')[1];
            if (!entries[master]) {
                entries[master] = [];
            }
            entries[master].push(name.replace('*.', ''));
        }
    });


    var changed = true;
    while (changed) {
        changed = false;
        for (var name in entries) {
            var current = entries[name];
            for (var sub in current) {
                var candidate = current[sub].split('.');
                var entryName = candidate[candidate.length - 1];
                if (entryName != name && entries[entryName] && !current['*' + entryName]) {
                    changed = true;
                    current['*' + candidate[candidate.length - 1]] = '*';
                    entries[entryName].forEach(function (itemToAdd) {
                        current.push(current[sub].replace('.' + entryName, '.' + itemToAdd));
                    });
                }
            }
        }
    }

    var new_state_tree = [];

    stateTree.forEach(function (name) {
        var parts = name.split('.');
        if (parts[0] != '*') {
            var tail = parts[parts.length - 1];
            new_state_tree.push(name);
            if (entries[tail]) {
                entries[tail].forEach(function (newEntry) {
                    new_state_tree.push(name.replace(tail, newEntry));
                });
            }
        }
    });

    return new_state_tree;
}

function createStateOptions(stateTree, stateConfigs) {
    return stateTree
    .map(function (name) {
        var hierarchy = name.split('.');

        // caso algum dos subestados nao exista, tiramos esse estado da reta
        var notFound = hierarchy.filter(function (part) { return !stateConfigs[part]; });
        if (notFound.length > 0)
            return null;

        var clone = JSON.parse(JSON.stringify(stateConfigs[hierarchy[hierarchy.length - 1]]));
        clone.name = name;

        // colocamos o layer referente ao nível hierárquico do state, caso o target não seja modal (modal é sempre modal)
        if (clone.target != 'modal' && clone.target != 'miolo' && clone.target != 'carrinho' && clone.target != 'filtros' && clone.target != 'correios' && clone.target != "entrega-agendada" && clone.target != "layer2")
            clone.target = hierarchy.length < 2 ? clone.target : 'layer' + (hierarchy.length - 1);

        return clone;
    })
    .filter(function (obj) { return !!obj; });
}

/*
	$routeProvider
		.when('/home', {
			templateUrl: 'templates/home.html'
		})
		.otherwise({
			redirectTo: '/home'
		}); 
  
	$locationProvider.hashPrefix('!');

	// This is required for Browser Sync to work poperly
	$httpProvider.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
}]);
*/

/*================================================================
=>                  client App Run()  
==================================================================*/


/* ---> Do not delete this comment (Values) <--- */

/* ---> Do not delete this comment (Constants) <--- */