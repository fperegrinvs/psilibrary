/// <reference path="../../Scripts/typings/angularjs/angular.d.ts" />
/// <reference path="../../Scripts/typings/angularjs/angular-resource.d.ts" />
/*================================================================
=>                  App = client
==================================================================*/
/*global angular*/
var Psilibrary;
(function (Psilibrary) {
    var app = angular.module('client', ["ngCookies", "ngResource", "ngSanitize", "ngRoute", "ngAnimate", "ui.utils", "ui.bootstrap", 
        "ui.router", "psilibrary.controllers", 'psilibrary.templates', 'psilibrary.config', 'psilibrary.services', 'psilibrary.directives', 
        'psilibrary.providers', 'frapontillo.bootstrap-duallistbox', 'templates', 'ngSanitize', 'textAngular']);
    angular.module("psilibrary.controllers", []);
    angular.module("psilibrary.templates", []);
    angular.module("psilibrary.services", []);
    angular.module("psilibrary.config", []);
    angular.module("psilibrary.directives", []);
    angular.module("psilibrary.providers", []);

    app.config(['showErrorsConfigProvider', function(showErrorsConfigProvider) {
      showErrorsConfigProvider.showSuccess(true);
    }]);

    function addSubStates(stateTree) {
        var entries = {};
        // mapeando dependencias
        stateTree.forEach(function (name) {
            if (name.split('.')[0] === '*') {
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
                    if (entryName !== name && entries[entryName] && !current['*' + entryName]) {
                        changed = true;
                        current['*' + candidate[candidate.length - 1]] = '*';
                        entries[entryName].forEach(function (itemToAdd) {
                            current.push(current[sub].replace('.' + entryName, '.' + itemToAdd));
                        });
                    }
                }
            }
        }
        var newStateTree = [];
        stateTree.forEach(function (name) {
            var parts = name.split('.');
            if (parts[0] !== '*') {
                var tail = parts[parts.length - 1];
                newStateTree.push(name);
                if (entries[tail]) {
                    entries[tail].forEach(function (newEntry) {
                        newStateTree.push(name.replace(tail, newEntry));
                    });
                }
            }
        });
        return newStateTree;
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
            return clone;
        })
            .filter(function (obj) { return !!obj; });
    }
    app.config([
        '$routeProvider', '$locationProvider', '$httpProvider', '$stateProvider', '$urlRouterProvider', function ($routeProvider, $locationProvider, $httpProvider, $stateProvider, $urlRouterProvider) {
            'use strict';
            //Enable cross domain calls
            $urlRouterProvider.otherwise('/404');
            var createState = function (name, url, partial, controllerName, view) {
                var views = {};
                views[view + "@"] = { templateUrl: "app/templates/" + partial + ".html", controller: controllerName + "Ctl" };
                return $stateProvider.state(name, { url: url, views: views });
            };
            var stateTree = [
                'home',
                'home2',
                'admin',
                'entryType',
                'entryTypeCreate',
                'entryTypeEdit',
                'category',
                'categoryCreate',
                'categoryEdit',
                '404',
                'entry',
                'entryCreate',
                'entryEdit',
                'facebookcallback',
                'search',
                'medline',
                'view',
            ];
            var stateConfigs = {
                'home': { url: '', partial: 'search', controller: 'search', target: 'miolo' },
                'home2': { url: '/', partial: 'search', controller: 'search', target: 'miolo' },
                'search': {url: '/search', partial: 'search', controller: 'search', target: 'miolo'},
                'admin': { url: '/admin', controller: 'admin', partial: 'admin', target: 'miolo'},
                'entryType': { url: '/entryType', controller: 'entryTypeList', partial: 'entryTypeList', target: 'miolo' },
                'entryTypeEdit': { url: '/entryType/{id}', controller: 'entryTypeEdit', partial: 'entryTypeEdit', target: 'miolo' },
                'entryTypeCreate': { url: '/entryType/new', controller: 'entryTypeCreate', partial: 'entryTypeEdit', target: 'miolo' },
                'category': { url: '/category', partial: 'categoryList', controller: 'categoryList', target: 'miolo' },
                'categoryCreate': { url: '/category/new', partial: 'categoryEdit', controller: 'categoryCreate', target: 'miolo' },
                'categoryEdit': { url: '/category/{id}', partial: 'categoryEdit', controller: 'categoryEdit', target: 'miolo' },
                '404': { url: '/404', partial: '404', controller: 'home', target: 'miolo' },
                'entry': { url: '/entry', partial: 'entryList', controller: 'entryList', target: 'miolo' },
                'entryCreate': { url: '/entry/new', partial: 'entryEdit', controller: 'entryCreate', target: 'miolo' },
                'entryEdit': { url: '/entry/{id}', partial: 'entryEdit', controller: 'entryEdit', target: 'miolo' },
                'facebookcallback' : { url: '/auth/facebook/callback', partial: 'home', controller: 'home', target: 'miolo'},
                'medline': {url: '/medline', partial: 'medline', controller: 'medline', target: 'miolo'},
                'view': {url: '/{id}', partial: 'view', controller: 'view', target: 'miolo'},
            };

            stateTree = addSubStates(stateTree);
            var stateoptions = createStateOptions(stateTree, stateConfigs);
            stateoptions.forEach(function (state) { createState(state.name, state.url, state.partial, state.controller, state.target); });
            // This is required for Browser Sync to work poperly 
            $httpProvider.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
        }
    ])
        .config([
        '$locationProvider', function ($location) {
            $location.hashPrefix('!');
        }
    ])
        .run([
        '$rootScope', '$state', '$stateParams', function ($rootScope, $state, $stateParams) {
            // It's very handy to add references to $state and $stateParams to the $rootScope
            // so that you can access them from any scope within your applications.For example,
            // <li ui-sref-active="active }"> will set the <li> // to active whenever
            // 'contacts.list' or one of its decendents is active.
            $rootScope.$state = $state;
            $rootScope.$stateParams = $stateParams;
            //$rootScope.path = window.path;
            $rootScope.$on('$stateChangeStart', function (event, toState, toParams, fromState, fromParams) {
                if (window.testing) {
                    console.log('$stateChangeStart to ' + toState.to + '- fired when the transition begins. toState,toParams : \n', toState, toParams);
                }
            });
            $rootScope.$on('$stateChangeError', function (event, toState, toParams, fromState, fromParams) {
                if (window.testing) {
                   console.log('$stateChangeError - fired when an error occurs during transition.');
                }
                //console.log(arguments);
            });
            $rootScope.$on('$stateChangeSuccess', function (event, toState, toParams, fromState, fromParams) {
                if (window.testing) {
                    console.log('$stateChangeSuccess to ' + toState.name + '- fired once the state transition is complete.');
                }
            });
        }
    ]); // helper functions
    function getScope(e) {
        return angular.element(e).scope();
    }
    Psilibrary.getScope = getScope;
    function getParentScope(e) {
        return angular.element(e).scope().$parent;
    }
    Psilibrary.getParentScope = getParentScope;
})(Psilibrary || (Psilibrary = {}));

var Utils;
(function (Utils) {
    /**
     * Register new controller.
     *
     * @param className
     * @param services
     */
    function registerController(className, services) {
        if (services === void 0) { services = []; }
        var controller = 'client.controllers.' + className;
        //services.push(app.controllers[className]);
        angular.module('client.controllers').controller(controller, services);
    }
    Utils.registerController = registerController;
    /**
     * Register new filter.
     *
     * @param className
     * @param services
     */
    function registerFilter(className, services) {
        if (services === void 0) { services = []; }
        var filter = className.toLowerCase();
        //services.push(() => (new app.filters[className]()).filter);
        angular.module('client.filters').filter(filter, services);
    }
    Utils.registerFilter = registerFilter;
    /**
     * Register new directive.
     *
     * @param className
     * @param services
     */
    function registerDirective(className, services) {
        if (services === void 0) { services = []; }
        var directive = className[0].toLowerCase() + className.slice(1);
        //services.push(() => new app.directives[className]());
        angular.module('client.directives').directive(directive, services);
    }
    Utils.registerDirective = registerDirective;
    /**
     * Register new service.
     *
     * @param className
     * @param services
     */
    function registerService(className, services) {
        if (services === void 0) { services = []; }
        var service = className[0].toLowerCase() + className.slice(1);
        //services.push(() => new app.services[className]());
        angular.module('app.services').factory(service, services);
    }
    Utils.registerService = registerService;
})(Utils || (Utils = {}));

function defaultPromise(h, $q){
    var deferred = $q.defer();
    h.success(function (data, status) {
        deferred.resolve(data, status);
    });
    h.error(function (data, status) {
        deferred.reject(data, status);
    });
    return deferred.promise;
}
