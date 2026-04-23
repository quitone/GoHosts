export namespace config {
	
	export class Scheme {
	    id: string;
	    name: string;
	    type: string;
	    content?: string;
	    url?: string;
	    enabled: boolean;
	    createdAt: string;
	    updatedAt: string;
	    refreshTime: number;
	
	    static createFrom(source: any = {}) {
	        return new Scheme(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.type = source["type"];
	        this.content = source["content"];
	        this.url = source["url"];
	        this.enabled = source["enabled"];
	        this.createdAt = source["createdAt"];
	        this.updatedAt = source["updatedAt"];
	        this.refreshTime = source["refreshTime"];
	    }
	}
	export class Config {
	    schemes: Scheme[];
	    activeScheme: string;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.schemes = this.convertValues(source["schemes"], Scheme);
	        this.activeScheme = source["activeScheme"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace runtime {
	
	export class EnvironmentInfo {
	    buildType: string;
	    platform: string;
	    arch: string;
	
	    static createFrom(source: any = {}) {
	        return new EnvironmentInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.buildType = source["buildType"];
	        this.platform = source["platform"];
	        this.arch = source["arch"];
	    }
	}

}

