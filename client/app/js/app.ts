/// <reference path="../../Scripts/typings/angularjs/angular.d.ts" />
/// <reference path="../../Scripts/typings/angularjs/angular-resource.d.ts" />
/*================================================================
=>                  App = client
==================================================================*/
/*global angular*/
module Psilibrary {
    var app = angular.module('client', ["ngCookies", "ngResource", "ngSanitize", "ngRoute", "ngAnimate", "ui.utils", "ui.bootstrap", "ui.router", "psilibrary.controllers", 'psilibrary.templates', 'psilibrary.config', 'psilibrary.services', 'psilibrary.directives']);

    angular.module("psilibrary.controllers", []);
    angular.module("psilibrary.templates", []);
    angular.module("psilibrary.services", []);
    angular.module("psilibrary.config", []);
    angular.module("psilibrary.directives", []);

    function addSubStates(stateTree) {
        var entries = {};

        // mapeando dependencias
        stateTree.forEach(name => {
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
                        entries[entryName].forEach(itemToAdd => {
                            current.push(current[sub].replace('.' + entryName, '.' + itemToAdd));
                        });
                    }
                }
            }
        }

        var newStateTree = [];

        stateTree.forEach(name => {
            var parts = name.split('.');
            if (parts[0] !== '*') {
                var tail = parts[parts.length - 1];
                newStateTree.push(name);
                if (entries[tail]) {
                    entries[tail].forEach(newEntry => {
                        newStateTree.push(name.replace(tail, newEntry));
                    });
                }
            }
        });

        return newStateTree;
    }

    function createStateOptions(stateTree, stateConfigs) {
        return stateTree
            .map(name => {
                var hierarchy = name.split('.');

                // caso algum dos subestados nao exista, tiramos esse estado da reta
                var notFound = hierarchy.filter(part => { return !stateConfigs[part]; });
                if (notFound.length > 0)
                    return null;

                var clone = JSON.parse(JSON.stringify(stateConfigs[hierarchy[hierarchy.length - 1]]));
                clone.name = name;

                return clone;
            })
            .filter(obj => { return !!obj; });
    }

    app.config([
            '$routeProvider', '$locationProvider', '$httpProvider', '$stateProvider', '$urlRouterProvider', ($routeProvider, $locationProvider, $httpProvider, $stateProvider, $urlRouterProvider) => {
                'use strict';
                //Enable cross domain calls

                $urlRouterProvider.otherwise('/404');


                var createState = (name, url, partial, controllerName, view) => {
                    var views = {};
                    views[view + "@"] = { templateUrl: "templates/" + partial + ".html", controller: controllerName + "Ctl" };
                    return $stateProvider.state(name, { url: url, views: views });
                };

                var stateTree = [
                    'home',
                    'home2',
                    'entryType',
                    'entryTypeCreate',
                    'entryTypeEdit',
                    'category',
                    'categoryCreate',
                    'categoryEdit',
                    '404',
                ];

                var stateConfigs = {
                    'home': { url: '', partial: 'home', controller: 'home', target: 'miolo' },
                    'home2': { url: '/', partial: 'home', controller: 'home', target: 'miolo' },
                    'entryType': { url: '/entryType', controller: 'entryTypeList', partial: 'entryTypeList', target: 'miolo' },
                    'entryTypeEdit': { url: '/entryType/{id}', controller: 'entryTypeEdit', partial: 'entryTypeEdit', target: 'miolo' },
                    'entryTypeCreate': { url: '/entryType/new', controller: 'entryTypeCreate', partial: 'entryTypeEdit', target: 'miolo' },
                    'category': { url: '/category', partial: 'categoryList', controller: 'categoryList', target: 'miolo' },
                    'categoryCreate': { url: '/category/new', partial: 'categoryEdit', controller: 'categoryCreate', target: 'miolo' },
                    'categoryEdit': { url: '/category/{id{', partial: 'categoryEdit', controller: 'categoryEdit', target: 'miolo' },                    
                    '404': { url: '/404', partial: '404', controller: 'home', target: 'miolo' },
                };

                stateTree = addSubStates(stateTree);
                var stateoptions = createStateOptions(stateTree, stateConfigs);

                stateoptions.forEach(state => { createState(state.name, state.url, state.partial, state.controller, state.target); });

                // This is required for Browser Sync to work poperly 
                $httpProvider.defaults.headers.common['X-Requested-With'] = 'XMLHttpRequest';
            }
        ])
        .config([
            '$locationProvider', $location => {
                $location.hashPrefix('!');
            }
        ])
        .run([
            '$rootScope', '$state', '$stateParams', ($rootScope, $state, $stateParams) => {
                // It's very handy to add references to $state and $stateParams to the $rootScope
                // so that you can access them from any scope within your applications.For example,
                // <li ui-sref-active="active }"> will set the <li> // to active whenever
                // 'contacts.list' or one of its decendents is active.
                $rootScope.$state = $state;
                $rootScope.$stateParams = $stateParams;
                //$rootScope.path = window.path;

                $rootScope.$on('$stateChangeStart', (event, toState, toParams, fromState, fromParams) => {
                    console.log('$stateChangeStart to ' + toState.to + '- fired when the transition begins. toState,toParams : \n', toState, toParams);
                });
                $rootScope.$on('$stateChangeError', (event, toState, toParams, fromState, fromParams) => {
                    console.log('$stateChangeError - fired when an error occurs during transition.');
                    //console.log(arguments);
                });
                $rootScope.$on('$stateChangeSuccess', (event, toState, toParams, fromState, fromParams) => {
                    console.log('$stateChangeSuccess to ' + toState.name + '- fired once the state transition is complete.');
                });
            }
        ]); // helper functions
    export function getScope(e) {
        return angular.element(e).scope();
    }

    export function getParentScope(e) {
        return angular.element(e).scope().$parent;
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
}

module Utils {
    export module controllers { }
    export module directives { }
    export module filters { }
    export module services { }

    export interface IController { }
    export interface IDirective {
        restrict: string;
        link($scope: ng.IScope, element: JQuery, attrs: ng.IAttributes): any;
    }
    export interface IFilter {
        filter(input: any, ...args: any[]): any;
    }
    export interface IService { }

    /**
     * Register new controller.
     *
     * @param className
     * @param services
     */
    export function registerController(className: string, services = []) {
        var controller = 'client.controllers.' + className;
        //services.push(app.controllers[className]);
        angular.module('client.controllers').controller(controller, services);
    }

    /**
     * Register new filter.
     *
     * @param className
     * @param services
     */
    export function registerFilter(className: string, services = []) {
        var filter = className.toLowerCase();
        //services.push(() => (new app.filters[className]()).filter);
        angular.module('client.filters').filter(filter, services);
    }

    /**
     * Register new directive.
     *
     * @param className
     * @param services
     */
    export function registerDirective(className: string, services = []) {
        var directive = className[0].toLowerCase() + className.slice(1);
        //services.push(() => new app.directives[className]());
        angular.module('client.directives').directive(directive, services);
    }

    /**
     * Register new service.
     *
     * @param className
     * @param services
     */
    export function registerService(className: string, services = []) {
        var service = className[0].toLowerCase() + className.slice(1);
        //services.push(() => new app.services[className]());
        angular.module('app.services').factory(service, services);
    }

}
